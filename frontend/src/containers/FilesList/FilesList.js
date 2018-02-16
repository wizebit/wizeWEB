import React, {Component} from 'react';
import {connect} from 'react-redux';
import axios from 'axios';
import classes from "./FilesList.css";
import {API_URL} from "../../shared/const";
import Spinner from '../../components/UI/Spinner/Spinner';
import Button from '../../components/UI/Button/Button';
import Aux from '../../hoc/Aux/Aux';
import Modal from '../../components/UI/Modal/Modal';

class Index extends Component {
    state = {
        files: null,
        loading: false,
        error: null,
        modalContent: null,
        transferTo: null
    };

    componentDidMount() {
        this.getFilesHandler();
    }

    getFilesHandler = () => {
        this.setState({loading: true});

        const config = {
            headers: {
                'Authorization': this.props.token
            }
        };

        axios.get(`${API_URL}/api/get-files-list`, config)
            .then(response => {
                this.setState({files: response.data.userFiles ? response.data.userFiles : [], loading: false}
            )})
            .catch(error => this.setState({error: error.response.data.message, loading: false}))
    };

    downloadFileHandler = (relativePath, filename) => {
        // const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = API_URL+relativePath;
        link.setAttribute('download', filename);
        document.body.appendChild(link);
        link.click();
    };

    onDeleteHandler = (filename) => {
        // console.log(filename);
        this.setState({loading: true});

        const config = {
            headers: {
                'Authorization': this.props.token
            }
        };

        axios.post(`${API_URL}/api/delete-file`, {filename: filename}, config)
            .then(response => {
                console.log(response.data.message);
                this.setState({loading: false});
                this.getFilesHandler();
                this.modalCloseHandler();
            })
            .catch(error => {
                this.setState({
                    error: error.response.data.message,
                    modalContent: <p>{error.response.data.message}</p>,
                    loading: false
                });
                this.modalCloseHandler();
            })
    };

    showDeleteModalHandler = (date, name) => {
        this.setState({
            modalContent: <div>
                <p>Are you sure?</p>
                <Button onClick={() => this.onDeleteHandler(date+"~"+name)}>Ok</Button>
            </div>
        });
    };

    onTransferHandler = (filename) => {
        console.log({filename: filename, transfer_to: this.state.transferTo});
        this.setState({loading: true});

        const config = {
            headers: {
                'Authorization': this.props.token
            }
        };

        axios.post(`${API_URL}/api/transfer-file`, {filename: filename, transferTo: this.state.transferTo}, config)
            .then(response => {
                console.log(response.data.message);
                this.setState({loading: false});
                this.getFilesHandler();
                this.modalCloseHandler();
            })
            .catch(error => {
                this.setState({
                    error: error.response.data.message,
                    modalContent: <p>{error.response.data.message}</p>,
                    loading: false
                });
                this.modalCloseHandler();
            })
    };

    showTransferModalHandler = (date, name) => {
        this.setState({
            modalContent: <div>
                <div>
                    <label htmlFor="transferTO">Enter public key of user, who will own this file.</label>
                    <input
                        type="text"
                        id="transferTO"
                        onChange={(e) => this.setState({transferTo: e.target.value})}
                    />
                </div>
                <Button onClick={() => this.onTransferHandler(date+"~"+name)}>Ok</Button>
            </div>
        });
    };


    modalCloseHandler = () => {
        this.setState({modalContent: null});
    };

    render() {
        let list = <Spinner />;

        if (Array.isArray(this.state.files)) {
            if (this.state.files.length !== 0) {
                list = <ul className={classes.FilesList}>
                    <li>
                        <span>Name</span>
                        <span>Upload Date</span>
                        <span>&nbsp;</span>
                        <span>&nbsp;</span>
                        <span>&nbsp;</span>
                    </li>
                    {
                        this.state.files.map((file, index) => <li key={index}>
                            <span>{file.name}</span>
                            <span>{new Date(file.uploadDate * 1000).toString()}</span>
                            <span>
                                <Button
                                    onClick={() => this.downloadFileHandler(file.relativePath, file.uploadDate+"~"+file.name)}
                                >
                                    Download
                                </Button>
                            </span>
                            <span>
                                <Button
                                    onClick={() => this.showDeleteModalHandler(file.uploadDate, file.name)}
                                >
                                    Delete
                                </Button>
                            </span>
                            <span>
                                <Button
                                    onClick={() => this.showTransferModalHandler(file.uploadDate, file.name)}
                                >
                                    Transfer file
                                </Button>
                        </span>
                        </li>)
                    }
                </ul>
            } else {
                list = <div>
                    <h2>
                        You don't have any files yet.
                    </h2>
                </div>
            }
        }

        return <Aux>
            <Modal show={ this.state.modalContent }
                   modalClosed={() => this.modalCloseHandler()}
            >
                { this.state.modalContent }
            </Modal>
            <div>
                <h1>Your files list</h1>

                {list}
            </div>
        </Aux>;
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.authKey,
        isAuth: state.auth.authKey !== null,
    }
};

export default connect(mapStateToProps)(Index);

// TODO: sharing