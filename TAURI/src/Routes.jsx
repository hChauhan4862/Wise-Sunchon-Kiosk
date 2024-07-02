import { Routes, Route, } from "react-router-dom";

import Home from "./pages/Home/Home";
import Floor from "./pages/Floor/Floor";

const Router = () => {

    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/floor/:id" element={<Floor />} />
        </Routes>
    )
}

export default Router