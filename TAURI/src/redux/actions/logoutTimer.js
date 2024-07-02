export const resetLogoutTimerDispatch = (time) => async (dispatch) => {
  dispatch({ type: "SET_LOGOUT_TIMER", data: time });
};