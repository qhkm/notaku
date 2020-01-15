import React, { useState } from "react";
import ReactDOM from "react-dom";
import {
    BrowserRouter as Router,
    Link,
    Switch,
    Route,
    useParams
} from "react-router-dom";
import axios from "axios";

const App = () => {
    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <Main />
                </Route>
                <Route exact path="/post/:id">
                    <Detail />
                </Route>
                <Route exact path="/test">
                    <Test />
                </Route>
            </Switch>
        </Router>
    );
};

const Main = () => {
    const [data, setData] = useState(null);
    const [title, setTitle] = useState("");
    const [body, setBody] = useState("");
    const [loading, setLoading] = useState(false);

    const handleTitleInput = e => {
        setTitle(e.target.value);
        console.log(title);
    };

    const handleBodyInput = e => {
        setBody(e.target.value);
        console.log(body);
    };

    const addData = async () => {
        const submitData = { title, body };
        console.log(submitData);
        let res = await axios.post(
            "http://localhost:8080/api/v1/posts/add",
            submitData
        );
        console.log(res);
    };

    const fetchData = async () => {
        const url = "http://localhost:8080/api/v1/posts";
        let res = await axios.get(url);
        setData(res.data.data);
        return res.data.data;
    };

    const deleteData = async idx => {
        const url = `http://localhost:8080/api/v1/post/${idx}`;
        let res = await axios.delete(url);
    };

    React.useEffect(() => {
        fetchData().then(data => {
            console.log(data);
        });
    }, []);
    return (
        <div>
            {loading && "loading"}
            <div>
                <Link to="/">Home</Link>
                <Link to="/test">Test</Link>
            </div>
            <div>
                <label>title</label>
                <input value={title} onChange={handleTitleInput} />
            </div>
            <div>
                <label>body</label>
                <input value={body} onChange={handleBodyInput} />
            </div>
            <button onClick={addData}>Add</button>
            {data &&
                data.map((el, idx) => (
                    <ul key={idx}>
                        <li>
                            <Link to={`/post/${el.id}`}>
                                <span>{el.title}</span>
                                <span>{el.body}</span>
                            </Link>
                            <button onClick={() => deleteData(el.id)}>
                                delete
                            </button>
                        </li>
                    </ul>
                ))}
        </div>
    );
};

const Detail = () => {
    const param = useParams();
    const [post, setPost] = useState(null);
    const [title, setTitle] = useState("");
    const [body, setBody] = useState("");

    const handleTitleInput = e => {
        setTitle(e.target.value);
        console.log(title);
    };

    const handleBodyInput = e => {
        setBody(e.target.value);
        console.log(body);
    };
    const fetchSinglePost = async () => {
        const res = await axios.get(
            `http://localhost:8080/api/v1/post/${param.id}`
        );
        setPost(res.data.data);
    };

    const updatePost = async () => {
        const updateData = { title, body };
        const res = await axios.put(
            `http://localhost:8080/api/v1/post/${param.id}`
        );
        console.log(res);
    };
    React.useEffect(() => {
        fetchSinglePost();
    }, []);

    console.log(post);

    return (
        <div>
            <h1>{post && post.title}</h1>
            <h3>{post && post.body}</h3>
            <div>
                <label>title</label>
                <input value={title} onChange={handleTitleInput} />
            </div>
            <div>
                <label>body</label>
                <input value={body} onChange={handleBodyInput} />
            </div>
            <button onClick={updatePost}>Update post</button>
        </div>
    );
};
const Test = () => <>test</>;

ReactDOM.render(<App />, document.getElementById("app"));
