import React, { useState } from "react";
import ReactDOM from "react-dom";
import axios from "axios";

const App = () => {
    const [data, setData] = useState(null);

    const [posts, setPosts] = useState([])
    const fetchData = async () => {
        const url = "http://localhost:8080/api/v1/posts";
        let res = await axios.get(url);
        setData(res.data.data);
    };
    console.log(data);

    React.useEffect(() => {
        let data = fetchData();
    }, []);

    return (
        <div>
            {data && data.map((el, idx) => <li key={idx}>{el.title}</li>)}
        </div>
    );
};

ReactDOM.render(<App />, document.getElementById("app"));
