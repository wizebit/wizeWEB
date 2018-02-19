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
        error: error.message
    }
};

export const cleanError = () => {
    return {
        type: actionTypes.AUTH_CLEAN_ERROR
    }
};

export const logout = () => {
    localStorage.removeItem('wise-bit-auth-key');
    localStorage.removeItem('wise-bit-auth-key-expiration-date');

  return {
      type: actionTypes.AUTH_LOGOUT
  }
};

export const auth = (publicKey, aesKey) => {
  return dispatch => {
      dispatch(authStart());
      axios.post(`${API_URL}/auth/sign-in`, {publicKey: publicKey, aesKey: aesKey}, {headers: {'Accept': 'application/json', 'Content-Type': 'application/json'}})
          .then(response => {
              const expirationDate = new Date(new Date().getTime() + response.data.expiresIn * 1000);
              localStorage.setItem('wise-bit-auth-key', response.data.accessToken);
              localStorage.setItem('wise-bit-auth-key-expiration-date', expirationDate);

              dispatch(checkAuthTimeout(response.data.expiresIn));

              dispatch(authSuccess(response.data));
          })
          .catch(error => {
              console.log(error.response.data);
              dispatch(authFail(error.response.data));
          });
  };
};

export const checkAuthTimeout = (expirationTime) => {
    return dispatch => {
        setTimeout(() => {dispatch(logout());}, expirationTime * 1000)
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
              dispatch(checkAuthTimeout(expirationDate.getTime() - new Date().getTime()));
              dispatch(authSuccess({accessToken: authKey}));
          }
      }
  }
};