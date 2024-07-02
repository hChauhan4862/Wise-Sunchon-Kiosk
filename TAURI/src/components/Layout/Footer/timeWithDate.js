export function timeInterval() {

  function updateCurrentTime() {
    const currentTime = new Date();
    const hours = currentTime.getHours();
    const minutes = currentTime.getMinutes();
    const ampm = hours >= 12 ? 'PM' : 'AM';

    // Format hours to 12-hour clock
    const formattedHours = hours;

    // Pad minutes with leading zero if necessary
    const formattedMinutes = minutes < 10 ? `0${minutes}` : minutes;

    const timeString = `${formattedHours}:${formattedMinutes}`;

    // Update the content of the <h6> element with the current time
    const timeElement = document.getElementById('hm_ampm');
    timeElement.textContent = timeString;
  }

  // Call the function once to display the initial time immediately
  updateCurrentTime();

  // Update the time every second (1000 milliseconds)
  // setInterval(updateCurrentTime, 1000)
}


export const TimerWithDate = () => {
  const time = new Date();
  const h = time.getHours();
  const mint = time.getMinutes();
  const s = time.getSeconds();
  const d = time.getDay();
  const dt = time.getDate();
  const m = time.getMonth();
  const y = time.getFullYear();
  const exch = h >= 12 ? "pm" : "am";

  const days = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
  ];
  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  let hours = h > 12 ? h - 12 : h;
  hours = hours < 10 ? "0" + hours : hours;
  const minutes = mint < 10 ? "0" + mint : mint;
  const seconds = s < 10 ? "0" + s : s;
  const day = days[d];
  const date = dt < 10 ? "0" + dt : dt;
  const month = months[m];
  const year = y;
  const ampm = exch;

  return { hours, minutes, seconds, day, date, month, year, ampm };
};