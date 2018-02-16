import React, {Component} from 'react';
import axios from 'axios';
import {connect} from 'react-redux';
import classes from './FileUpload.css';
import Dropzone from 'react-dropzone';
import {API_URL} from "../../shared/const";

class Files extends Component {
    state = {
        progress: 0,
        loading: false,
        error: false
    };

    onDropHandler = (file) => {
        this.setState({loading: true});

        let data = new FormData();
        data.append('file', file);

        const config = {
            onUploadProgress: progressEvent => {
                let percentCompleted = Math.floor((progressEvent.loaded * 100) / progressEvent.total);
                // do whatever you like with the percentage complete
                // maybe dispatch an action that will update a progress bar or something
                setTimeout(() => this.setState({progress: percentCompleted}), 400);
            },
            headers: {
                'Authorization': this.props.token,
                'Content-Type': 'multipart/form-data'
            }
        };

        axios.put(`${API_URL}/api/upload-file`, data, config)
            .then(res => {
                console.log(res.data);
                this.setState({
                    progress: 0,
                    loading: false,
                    error: false
                });
            })
            .catch(err => {
                console.log(err.message);
                this.setState({
                    progress: 0,
                    loading: false,
                    error: err.message
                });
            });
    };

    render() {
        let progress = null;
        if (this.state.loading) {
            progress = <div className={classes.ProgressBar}>
                <div className={classes.ProgressLine} style={{width: this.state.progress+"%"}} />
                <div className={classes.Percentage}>
                    {this.state.progress !== 100 ? this.state.progress+"%" : "Sending..."}
                </div>
            </div>;
        }

        return <div>
            <h1>Files page</h1>
            <Dropzone
                onDrop={files => this.onDropHandler(files[0])}
                // className={classes.Dropzone}
                //
                //  or
                //
                // style={{
                //     width: "100%",
                //     height: "50vh",
                //     padding: "30px",
                //     boxSizing: "border-box",
                //     cursor: "pointer",
                //     display: "flex",
                //     justifyContent: "center",
                //     alignItems: "center"
                // }}
            >
                <p>Try dropping some files here, or click to select files to upload.</p>
            </Dropzone>

            {progress}
        </div>
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.authKey,
        isAuth: state.auth.authKey !== null,
    }
};

export default connect(mapStateToProps)(Files);