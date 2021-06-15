import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Button } from '@material-ui/core';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
            Something goes here hello world jkahdfklhasasdfsfdf
        </a>
      </header>
      <Button color="primary">Hello World I am here</Button>;
    </div>
  );
}

export default App;
