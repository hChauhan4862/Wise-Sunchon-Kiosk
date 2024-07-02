import { combineReducers } from 'redux';

import authReducer from './auth';
import seatBookingReducer from './seatBooking';
import logoutTimerReducer from './logoutTimerReducer';
import setupReducer from './setup';
import seatsReducer from './seats';

export default combineReducers({
    authReducer,
    seatBookingReducer,
    logoutTimerReducer,
    setupReducer,
    seatsReducer
});
