package astore_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/enfabrica/enkit/astore/client/astore"
	rpcAstore "github.com/enfabrica/enkit/astore/rpc/astore"
	astoreServer "github.com/enfabrica/enkit/astore/server/astore"
	"github.com/enfabrica/enkit/lib/client/ccontext"
	"github.com/enfabrica/enkit/lib/khttp/ktest"
	"github.com/enfabrica/enkit/lib/logger"
	"github.com/enfabrica/enkit/lib/progress"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
)

func TestSid(t *testing.T) {
	rng := rand.New(rand.NewSource(0))
	for i := 0; i < 1000; i++ {
		value, err := astoreServer.GenerateSid(rng)
		assert.Nil(t, err)
		assert.Equal(t, 34, len(value), "value: %s", value)
	}
}

func TestUid(t *testing.T) {
	rng := rand.New(rand.NewSource(0))
	for i := 0; i < 1000; i++ {
		value, err := astoreServer.GenerateUid(rng)
		assert.Nil(t, err)
		assert.Equal(t, 32, len(value), "value: %s", value)
		assert.True(t, astore.IsUid(value))
	}
}

// TODO(aaahrens): fix client so that it's signed urls can depend on an interface for actual e2e testing
func TestServer(t *testing.T) {
	astoreDescriptor, killFuncs, err := ktest.RunAStoreServer()
	if killFuncs != nil {
		defer killFuncs.KillAll()
	}
	if err != nil {
		t.Fatal(err.Error())
	}
	// running this as test ping feature
	client := astore.New(astoreDescriptor.Connection)
	res, _, err := client.List("/test", astore.ListOptions{})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("list response is +%v \n", res)
	b, err := ioutil.ReadFile("./testdata/example.yaml")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("bytes are ", string(b))
	uploadFiles := []astore.FileToUpload{
		{Local: "./testdata/example.yaml"},
	}

	ctxWithLogger := ccontext.DefaultContext()
	ctxWithLogger.Logger = logger.DefaultLogger{Printer: log.Printf}
	ctxWithLogger.Progress = progress.NewDiscard

	uploadOption := astore.UploadOptions{
		Context: ctxWithLogger,
	}
	u, err := client.Upload(uploadFiles, uploadOption)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("erroring in client upload")
	}

	fmt.Printf("upload is +%v \n", u)
	storeResponse, err := astoreDescriptor.Server.Store(context.Background(), &rpcAstore.StoreRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if storeResponse.GetSid() == "" || storeResponse.GetUrl() == "" {
		t.Fatal(errors.New("invalid store response"))
	}
	resp, err := astoreDescriptor.Server.Commit(context.Background(), &rpcAstore.CommitRequest{
		Sid:          storeResponse.GetSid(),
		Architecture: "dwarvenx99",
		Path:         "127.0.0.1:9000/hello/work/example.yaml",
		Note:         "note",
		Tag:          []string{"something"},
	})
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println("finalizzing +%v", resp.Artifact)

}
