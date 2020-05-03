import React, { Component } from 'react';
import { Link } from 'react-router-dom'
import ReactModal from "react-modal";
import TweetModal from './TweetModal.js'
import profilePic from '../img/user-image-with-black-background.png'

export class Navbar extends Component {
	constructor() {
		super();
		this.state = {
			showModal: false
		};
	}

	toggleModal = () => this.setState({showModal: !this.state.showModal});

	closeModal = () => this.setState({ showModal: false });

    render() {
        return (
            <div className="nav-container">
			 <nav className="main-menu">
			   <ul>
			     <Link to="/" style={{ textDecoration: 'none', color: '#1e91da' }}><li className="main-menu-item"><div className="rounder">Home</div></li></Link>
			     <Link to="/explore" style={{ textDecoration: 'none', color: '#1e91da' }}><li className="main-menu-item">Hashtag</li></Link>
			     <Link to="/notifications" style={{ textDecoration: 'none', color: '#1e91da'}}><li className="main-menu-item">Notify</li></Link>
			     <Link to="/messages" style={{ textDecoration: 'none', color: '#1e91da'}}><li className="main-menu-item">Messages</li></Link>
			     <li><button id="btn" onClick={this.toggleModal}>
					Search Twitter
				</button></li>
				<ReactModal
							appElement={document.getElementById("root")}
							isOpen={this.state.showModal}
							contentLabel="onRequestClose Example"
							onRequestClose={this.closeModal}
							className="Modal"

						>
							<TweetModal />
				</ReactModal> 
				<li className="main-menu-item"><img className="profilePic" src={profilePic} alt="dsada" onClick={this.toggleModal}/></li>
			   </ul>
			 </nav>
			</div>
        );
    }
}

export default Navbar;