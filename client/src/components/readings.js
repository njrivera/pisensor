import React from 'react';
import Table from '@material-ui/core/Table';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import axios from 'axios';
import Button from '@material-ui/core/Button';
import TempReadingTimePicker from './timepicker';
import MuiPickersUtilsProvider from 'material-ui-pickers/MuiPickersUtilsProvider';
import MomentUtils from 'material-ui-pickers/utils/moment-utils'
import TextField from '@material-ui/core/TextField';
import { Paper } from '@material-ui/core';
import TempChart from './tempchart';

export default class ReadingsTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            rows: [],
            tempData: [],
            liveData: [],
            liveRange: 10,
            serial: '',
            serialList: '',
            start: new Date().toLocaleString(),
            end: new Date().toLocaleString()
        }

        this.ws = null;
        this.getReadings = this.getReadings.bind(this);
        this.setStartTime = this.setStartTime.bind(this);
        this.setEndTime = this.setEndTime.bind(this);
        this.sendFilter = this.sendFilter.bind(this);

        this.ws = new WebSocket("ws://localhost:5556/livetemp");
        this.ws.onmessage = msg => {
            let t = JSON.parse(msg.data);
            let temps = JSON.parse(JSON.stringify(this.state.liveData));

            if (t.serial === 'a') {
                temps.push({Time: new Date(t.time).toLocaleString(), Temperature: t.temp});
            }

            if (temps.length > this.state.liveRange) {
                temps.shift(1);
            }

            this.setState({liveData: temps});
        }
    }

    setStartTime(t) {
        this.setState({start: t});
    }

    setEndTime(t) {
        this.setState({end: t});
    }

    createData(id, time, serial, model, temp, unit) {
        return {id, time, serial, model, temp, unit};
    }

    getReadings() {
        let url = new URL('http://localhost:3000/api/readings/betweentimes');

        url.searchParams.set('serial', this.state.serial);
        url.searchParams.set('start', this.state.start);
        url.searchParams.set('end', this.state.end);

        axios.get(url)
            .then(resp => {
                let respRows = [];
                let temps = [];
                let id = 0;
                resp.data.forEach(reading => {
                    id++
                    let t = new Date(reading.time).toLocaleString();
                    respRows.push(this.createData(id, t, this.state.serial, reading.model, reading.temp, reading.unit));
                    temps.push({Time: t, Temperature: reading.temp});
                });

                this.setState({rows: respRows});
                this.setState({tempData: temps});
            });
    }

    sendFilter(value) {
        if (this.ws) {
            let msg = {serials: this.state.serialList.split(',')};
            this.ws.send(JSON.stringify(msg));
        }
    }

    render() {
        return (
            <div>
                <br/>
                <TextField value={this.state.serialList} label="Serial List Filter" onChange={event => this.setState({serialList: event.target.value})}/>
                <Button variant="contained" color="primary" onClick={this.sendFilter}> Get Live Temp Data </Button>
                <Paper>
                    <TempChart tempData={this.state.liveData}/>
                </Paper>
                <br/>
                <TextField value={this.state.serial} label="Serial" onChange={event => this.setState({serial: event.target.value})}/>
                <MuiPickersUtilsProvider utils={MomentUtils}>
                    <TempReadingTimePicker setTime={this.setStartTime} time={this.state.start}/>
                    <TempReadingTimePicker setTime={this.setEndTime} time={this.state.end}/>
                </MuiPickersUtilsProvider>
                <br/>
                <Paper>
                    <TempChart tempData={this.state.tempData}/>
                </Paper>
                <Button variant="contained" color="primary" onClick={this.getReadings}> Get Historical Temp Readings </Button>
                <Paper>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Time</TableCell>
                                <TableCell>Serial</TableCell>
                                <TableCell>Model</TableCell>
                                <TableCell>Temperature</TableCell>
                                <TableCell>Unit</TableCell>
                            </TableRow>
                        </TableHead>

                        <TableBody>
                            {this.state.rows.map(row => {
                                return (
                                    <TableRow key={row.id}>
                                        <TableCell>{row.time}</TableCell>
                                        <TableCell>{row.serial}</TableCell>
                                        <TableCell>{row.model}</TableCell>
                                        <TableCell>{row.temp}</TableCell>
                                        <TableCell>{row.unit}</TableCell>
                                    </TableRow>
                                );
                            })}
                        </TableBody>
                    </Table>
                </Paper>
            </div>
        );
    }
}