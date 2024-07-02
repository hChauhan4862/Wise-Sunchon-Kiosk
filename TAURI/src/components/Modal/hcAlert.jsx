import Swal from "sweetalert2";
import store from "../../redux";
import withReactContent from "sweetalert2-react-content";
import { logoutUser } from "../../redux/actions/auth";

export const hcAlert = (type, message, logout) => {
    // type = 'Opps! | 'Success'
    const timers = store.getState().setupReducer.timers;
    const MySwal = withReactContent(Swal);
    MySwal.fire({
        title: type,
        showConfirmButton: false,
        width: 600,
        imageUrl: type == 'Opps!' ? '/images/cancel.png' : '/images/check.png',
        timer:    type == 'Opps!' ? (timers?.error || 1) * 1000 : (timers?.success || 1) * 1000,
        imageWidth: 110,
        html: <div>{(message)}</div>,
        allowOutsideClick: false,
      }).then((result) => {
        if (result.dismiss === Swal.DismissReason.timer && logout) {
            store.dispatch(logoutUser())
        }
    });
}
