import React, {Component} from 'react';
import axios from 'axios';
import {connect} from 'react-redux';
import classes from './WalletsList.css';
import {API_URL} from "../../shared/const";
import Spinner from '../../components/UI/Spinner/Spinner';

class WalletsList extends Component {
    state = {
        walletsList: null,
        error: null,
        loading: false
    };

    componentDidMount() {
        this.getWalletsListHandler();
    }

    getWalletsListHandler = () => {
        this.setState({loading: true});

        const config = {
            headers: {
                'X-ACCESS-TOKEN': this.props.token
            }
        };

        axios.get(`${API_URL}/api/get-wallets-list`, config)
            .then(response => {
                this.setState({walletsList: response.data.walletsList, loading: false}
                )})
            .catch(error => this.setState({error: error.response.data.message, loading: false}))
    };

    render() {
        let content;

        if (this.state.walletsList) {
            content = <ul>
                {this.state.walletsList.map((wallet, index) => <li key={index}>{wallet}</li>)}
            </ul>;
        }

        if (this.state.loading) {
            content = <Spinner />
        }

        return <div>
            <h1>Wallets List</h1>

            {content}
        </div>
    }
}

const mapStateToProps = state => {
    return {
        token: state.auth.authKey,
        isAuth: state.auth.authKey !== null,
    }
};

export default connect(mapStateToProps)(WalletsList);