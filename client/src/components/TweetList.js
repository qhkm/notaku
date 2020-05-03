import React from "react";
import { Route, Link } from "react-router-dom";
import axios from 'axios'
import ReactModal from "react-modal";
import Tweet from './Tweet.js'
import TweetModal from './TweetModal.js'
import profilePic from '../img/user-image-with-black-background.png'
import uploadImagePlaceholder from '../img/imagePlaceholder.png'
import gifPlaceholder from  '../img/gif.png'

class TweetList extends React.Component {
	state = {
		showModal: false,
		posts: []
	}

	componentDidMount(){
		axios.get('http://localhost:8080/api/v1/posts')
			.then(res => { 
				const posts = res.data.data; 
				this.setState({posts: posts})
			}).then(() => console.log(this.state.posts))
			.catch(err => console.log(err))
	}

	toggleModal = () => this.setState({showModal: !this.state.showModal});

	closeModal = () => this.setState({ showModal: false });

	render() {
		return (
			<div className="container">
				<div className="space" />
				<main className="content-area">
					<div className="card-top">
						<img className="profilePic" src={profilePic} alt="dsada" onClick={this.toggleModal}/>
						<button id="btn" onClick={this.toggleModal}>
							Whats happening?
						</button>
						<ReactModal
							appElement={document.getElementById("root")}
							isOpen={this.state.showModal}
							contentLabel="onRequestClose Example"
							onRequestClose={this.closeModal}
							className="Modal"

						>
							<TweetModal />
						</ReactModal>
						<img className="uploadImagePlaceholder" src={uploadImagePlaceholder} alt="dsada" onClick={this.toggleModal}/>
						<img className="gifPlaceholder" src={gifPlaceholder} alt="dsada" onClick={this.toggleModal}/>
					</div>
					{this.state.posts.map(({ id, body }) => (
						<li key={id} style={{ listStyleType: "none" }}>
							<Link
								to={{
									pathname: `/posts/${id}`,
									state: {
										 body
									}
								}}
								style={{
									textDecoration: "none",
									color: "grey"
								}}
							>
								<div className="card">{body}</div>
							</Link>
							
						</li>

					))}
				</main>
				<aside className="sidebar">
					<div className="sidebar-container">
						<h3>Sidebar</h3>
						<ul>
							<li>Items</li>
							<li>Are</li>
							<li>Listed</li>
							<li>Here</li>
						</ul>
					</div>

					<div className="sidebar-container">
						<h3>Sidebar</h3>
						<ul>
							<li>Items</li>
							<li>Are</li>
							<li>Listed</li>
							<li>Here</li>
						</ul>
					</div>
				</aside>
				<Route path={`/posts/:id`} component={Tweet} />
			</div>
		);
	}
}


export default TweetList;
