import * as actionTypes from './actionTypes';
import axios from 'axios';
import { API_URL } from '../../shared/const';

export const authStart = () => {
    return {
        type: actionTypes.AUTH_START
    }
};

export const authSuccess = (authData) => {
    return {
        type: actionTypes.AUTH_SUCCESS,
        authData: authData
    }
};

export const authFail = (error) => {
    return {
        type: actionTypes.AUTH_FAIL,
        error: error
    }
};

export const logout = () => {
    localStorage.removeItem('wise-bit-auth-key');
    localStorage.removeItem('wise-bit-auth-key-expiration-date');

  return {
      type: actionTypes.AUTH_LOGOUT
  }
};

export const auth = (privateKey) => {
  return dispatch => {
      dispatch(authStart());
      axios.post(`${API_URL}/auth/sign-in`, {private_key: privateKey}, {headers: {'Accept': 'application/json', 'Content-Type': 'application/json'}})
          .then(response => {
              const expirationDate = new Date(new Date().getTime() + response.data.expires_in * 1000);
              localStorage.setItem('wise-bit-auth-key', response.data.auth_key);
              localStorage.setItem('wise-bit-auth-key-expiration-date', expirationDate);

              dispatch(authSuccess(response.data));
              dispatch(checkAuthTimeout(response.data.expires_in));
          })
          .catch(error => {
              console.log(error.response);
              dispatch(authFail(error));
          });
  };
};

export const checkAuthTimeout = (expirationTime) => {
    return dispatch => {
        setTimeout(() => {
            dispatch(logout());
        }, expirationTime * 1000)
    }
};

export const authCheckState = () => {
  return dispatch => {
      const authKey = localStorage.getItem('wise-bit-auth-key');


      if (!authKey) {
          dispatch(logout());
      } else {
          const expirationDate = new Date(localStorage.getItem('wise-bit-auth-key-expiration-date'));
          if (expirationDate <= new Date()) {
              dispatch(logout());
          } else {
              dispatch(authSuccess({auth_key: authKey}));
              dispatch(checkAuthTimeout(expirationDate.getTime() - new Date().getTime()));
          }
      }
  }
};