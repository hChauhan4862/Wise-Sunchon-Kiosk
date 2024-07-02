const seatsReducer = (state = {}, action) => {
    switch (action.type) {
        case 'SET_SeatsListBySector':
            return { ...state, seatsListBySector: action.data }
        case 'ADD_SeatsMapUrl':
            let ARR = state?.seatsMapUrls || {}
            // Append To Object
            Object.assign(ARR, {[action.data.url]:action.data.data})
            // console.log(ARR, "##################################################")
            return { ...state, seatsMapUrls: ARR }
        default:
            return state;
    }
}

export default seatsReducer;