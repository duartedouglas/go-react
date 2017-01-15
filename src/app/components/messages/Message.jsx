import React, {Component} from 'react';
import fecha from "fecha";

class Message extends Component{
    render(){
        const {message} = this.props;
        let createdAt = fecha.format(fecha.parse(message.createdAt, "YYYY-MM-DD hh:mm:ss.SSS"), "DD/MM HH:MM::SS" );
        return (
            <li className="message">
                <div className="author">
                    <strong>{message.author}</strong>
                    <i className="timestamp">{createdAt}</i>
                </div>
                <div className="body">{message.body}</div>
            </li>
        )
    }
}

Message.propTypes = {
    message: React.PropTypes.object.isRequired
};

export default Message