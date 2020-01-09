import React from 'react'
import ReactDOM from 'react-dom'
import axios from 'axios'

const App = () => {

    const fetchData = async () => {
        const url = 'http://localhost:8080/api/v1/posts'
        let res = await axios.get(url)
        console.log(res)
    }

    React.useEffect(() => {
        let data = fetchData();
    })

    return (
        <div>test</div>
    )
}

ReactDOM.render(<App />, document.getElementById('app'))