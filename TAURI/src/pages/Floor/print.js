
import axios from "axios";
import jwt_decode from "jwt-decode";

export const printHandler = async (bookingdata, userData) => {
    // const userData = jwt_decode(getToken)

    console.log("nananan", userData.barcode)
    const useEndTime = bookingdata.use_end_time; // Assuming use_end_time is in the format "HH:MM:SS"

    // Convert useEndTime to a Date object
    const endTime = new Date(`2000-01-01T${useEndTime}`);

    // Calculate the extension time by subtracting one hour from the use end time
    endTime.setHours(endTime.getHours() - 1);

    // Format the extension time and use end time in the desired format
    const extensionTime = endTime.toTimeString().slice(0, 8); // Extract the time portion (HH:MM:SS)
    const formattedEndTime = useEndTime;
    try {
        const data = {
            data: [
                {
                    type: "text",
                    value: "SEAT BOOKING RECEIPT",
                    style: "text-align:center;",
                    css: {
                        "font-weight": "600",
                        "font-size": "15px",
                        "margin-bottom": "10px",
                        "margin-top": "3px",
                    },
                },
                {
                    type: "text",
                    value: "SUNCHON NATIONAL UNIVERSITY",
                    style: "text-align:center;",
                    css: {
                        "font-weight": "400",
                        "font-size": "10px",
                        "margin-bottom": "10px",
                        "margin-top": "3px",
                    },
                },
                {
                    type: "text",
                    value: `Booking No:  ${bookingdata.BookingId}`,
                    style: "text-align:left;color: black;",
                    css: { "font-size": "15px" },
                },
                // {
                //     type: "text",
                //     value: `Seat Location: ${bookingdata.floorName} / ${bookingdata.zoneName}  `,
                //     style: "text-align:left;color: black;",
                //     css: { "font-size": "15px" },
                // },
                {
                    type: "text",
                    value: `Seat No: ${bookingdata.seatNo}`,
                    style: "text-align:left;color: black;",
                    css: { "font-size": "15px" },
                },
                // {
                //     type: "text",
                //     value: `Seat Name: ${bookingdata.seatCategoryName}`,
                //     style: "text-align:left;color: black;",
                //     css: { "font-size": "15px" },
                // },
                {
                    type: "text",
                    value: `Number of Extentions: ${bookingdata.extension_count} Used |  3 Allowed`,
                    style: "text-align:left;color: black;",
                    css: { "font-size": "15px" },
                },
                {
                    type: "text",
                    value: `Name: ${bookingdata.userName} `,
                    style: "text-align:left;color: black;",
                    css: { "font-size": "15px" },
                },
                {
                    type: "text",
                    value: `Check in Time : ${bookingdata.use_start_time}`,
                    style: "text-align:left;color: black;",
                    css: { "font-size": "15px" },
                },
                {
                    type: "text",
                    value: `Check Out Time : ${bookingdata.use_end_time}`,
                    style: "text-align:left;color: black;",
                    css: { "font-size": "15px" },
                },
                {
                    type: "text",
                    value: `Extension Time: ${extensionTime} ~ ${formattedEndTime}`,
                    style: "text-align:left;color: black;",
                    css: { "font-size": "15px" },
                },
                {
                    // type: "barCode",
                    type: "qrCode",
                    value: userData?.barcode,
                    // height: 40, //barcode
                    // width: 1, //barcode
                    height: 120,
                    width: 120,
                    style: "margin: 10px 20px 20px 20px",
                    position: "center",

                    displayValue: true, // Display value below barcode,
                    fontsize: 12, //barcode,
                },
                {
                    type: "text",
                    value: "*If you return the seat when assignment when leaving",
                    style: "text-align:left;color: black;",
                    css: { "font-size": "12px" },
                },
                {
                    type: "text",
                    value: "*The rooom it will be helpful to other people to use it.",
                    style: "text-align:left;color: black;",
                    css: { "font-size": "12px" },
                },
            ],
        };

        const res = await axios.post("http://localhost:8084/kiosk/print", data);
        console.log(res)
        return res
     
    } catch (error) {
        console.log(error);
    }
};