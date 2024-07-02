
const seatBookingReducer =  (state = {data:null}, action) => {
    switch (action.type) {
        case 'SET_SEAT_BOOKING':
            return { ...state, seatBooking: action.data }
        default:
            return state;
    }
}

export default seatBookingReducer