import { resetLogoutTimerDispatch } from "../../redux/actions/logoutTimer";
import { logoutUser } from "../../redux/actions/auth";
import store from "../../redux";

export const SessionTimer = (time) => {
    time = time || new Date().getTime()
    store.dispatch(resetLogoutTimerDispatch(time))
}

export const SessionInterval = (handleLogout) => {
    SessionTimer()
    const timers = store.getState().setupReducer.timers
    const element = document.getElementById('time')
    element.style.display = 'block'

    let minutes = Math.floor(timers.floorPage / 60).toString().padStart(2, '0');
    let seconds = Math.floor(timers.floorPage % 60).toString().padStart(2, '0');
    element.innerHTML = `${minutes}:${seconds}`

    const interval = setInterval(() => {
        if (store.getState().authReducer.userData === null) {
            handleLogout('/')
            clearInterval(interval)
            element.style.display = 'none'
        }
        const time = (1000 * timers.floorPage) - (new Date().getTime() - parseInt(store.getState().logoutTimerReducer.resetLogoutTimerDispatch))
        const timeInSeconds = time / 1000;
        // const element = document.getElementById('time')
        // element.style.display = 'block'
        // element.
        if (timeInSeconds > 5) {
            element.style.boxShadow = 'rgb(95, 255, 95) 0px 10px 35px -12px inset, rgb(95, 255, 95) 0px 18px 36px -18px inset'
        } else {
            element.style.boxShadow = 'rgb(255, 95, 95) 0px 10px 35px -12px inset, rgb(255, 95, 95) 0px 18px 36px -18px inset'
        }
        setTimeout(() => {
            element.style.boxShadow = 'none'
        }, 500)
        minutes = Math.floor(timeInSeconds / 60).toString().padStart(2, '0');
        seconds = Math.floor(timeInSeconds % 60).toString().padStart(2, '0');
        element.innerHTML = `${minutes}:${seconds}`


        if ((new Date()).getTime() - parseInt(store.getState().logoutTimerReducer.resetLogoutTimerDispatch) > 1000 * timers.floorPage) {
            if (store.getState().authReducer.userData != null) {
                store.dispatch(logoutUser(handleLogout))
                SessionTimer(new Date().getTime() - 1000 * timers.floorPage)
                clearInterval(interval)
                element.style.display = 'none'
            }
        }
    }, 1000)
}

export const keyboardHideInterval = (handleLogout) => {
    SessionTimer()
    const timers = store.getState().setupReducer.timers
    const interval = setInterval(() => {
        if (store.getState().authReducer.userData === null) {
            handleLogout('/')
            clearInterval(interval)
        }
        if ((new Date()).getTime() - parseInt(store.getState().logoutTimerReducer.resetLogoutTimerDispatch) > 1000 * timers.floorPage) {
            if (store.getState().authReducer.userData != null) {
                store.dispatch(logoutUser(handleLogout))
                // handleLogout('/')
                SessionTimer(new Date().getTime() - 1000 * timers.floorPage)
                clearInterval(interval)
            }
        }
    }, 1000)
}