import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import NavBar from './components/navbar'
import ReadingsTable from './components/readings';
import MuiPickersUtilsProvider from 'material-ui-pickers/MuiPickersUtilsProvider';
import DateFnsUtils from 'material-ui-pickers/utils/date-fns-utils';

class App extends Component {
  render() {
    return (
      <div className="App">
        <MuiPickersUtilsProvider utils={DateFnsUtils}>
          <NavBar />
          <ReadingsTable />
        </MuiPickersUtilsProvider>
      </div>
    );
  }
}

export default App;
