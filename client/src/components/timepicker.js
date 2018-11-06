import React, { PureComponent } from 'react';
import DateTimePicker from 'material-ui-pickers/DateTimePicker';

export default class TempReadingTimePicker extends PureComponent {
    handleDateChange = (date) => {
        date = new Date(date).toLocaleString();
        this.props.setTime(date);
    }

    render() {
        return (
            <DateTimePicker
                onChange={this.handleDateChange}
                value={this.props.time}
            />
        );
    }
}