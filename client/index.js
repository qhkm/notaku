import React, { useState } from 'react'
import ReactDOM from 'react-dom'
import axios from 'axios'

const App = () => {

    const [posts, setPosts] = useState([])
    const fetchData = async () => {
        const url = 'http://localhost:8080/api/v1/posts'
        let res = await axios.get(url)
        setPosts(res.data.data)
    }
    console.log(posts)

    React.useEffect(() => {
        let data = fetchData();
    }, [])

    return (
        <ul>{posts !== null && posts.length > 0 &&
            posts.map((data, idx) => (<li key={idx}>{data.title}</li>
            ))}
        </ul>
    )
}

ReactDOM.render(<App />, document.getElementById('app'))