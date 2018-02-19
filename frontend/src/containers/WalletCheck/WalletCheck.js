import React, {Component} from 'react';
import axios from 'axios';
import {connect} from 'react-redux';
import classes from './WalletCheck.css';
import {API_URL} from "../../shared/const";
import Spinner from '../../components/UI/Spinner/Spinner';
import Input from '../../components/UI/Input/Input';

class WalletCheck extends Component {
    state = {
        inputVal: null,
        walletInfo: null,
        error: null,
        loading: false
    };

    onInputNumberHandler = (e) => {
        const trigger = e.target.value;
        this.setState({inputVal: trigger, walletInfo: null, loading: trigger, error: null});
        setTimeout(() => this.checkWalletHandler(trigger), 2000);
    };

    checkWalletHandler = (val) => {
        if (val === this.state.inputVal && this.state.inputVal) {
            const config = {
                headers: {
                    'X-ACCESS-TOKEN': this.props.token
                }
            };

            axios.get(`${API_URL}/api/wallet/${val}`, config)
                .then(response => {
                    this.setState({walletInfo: response.data.walletInfo, loading: false}
                    )})
                .catch(error => this.setState({error: error.response.data.message, loading: false}))
        }
    };

    render() {
        let content;

        if (this.state.walletInfo) {
            content = <div>
                {/*{ [].forEach.call(Object.keys(this.state.walletInfo), (item) => {return <p key={item}>{item}: {this.state.walletInfo[item]}</p>})}*/}
                <p>Credit: {this.state.walletInfo.credit}</p>
                <p>Success: {this.state.walletInfo.success ? "true" : "false"}</p>
            </div>;
        }

        if (this.state.error) {
            content = <div>
                Error: {this.state.error}
            </div>
        }

        if (this.state.loading) {
            content = <Spinner />
        }

        return <div>
            <h1>Wallet Check</h1>
            <div>
                <label>
                    <p>Enter wallet number bellow</p>
                    <Input changed={e => this.onInputNumberHandler(e)}/>
                </label>
            </div>

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

export default connect(mapStateToProps)(WalletCheck);