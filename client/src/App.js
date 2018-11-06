import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import NavBar from './components/navbar'
import ReadingsTable from './components/readings';

class App extends Component {
  render() {
    return (
      <div className="App">
        <NavBar />
        <ReadingsTable />
      </div>
    );
  }
}

export default App;
