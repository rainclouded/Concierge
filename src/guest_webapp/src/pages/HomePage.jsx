import React from 'react'
import { Link } from 'react-router-dom'

const HomePage = () => {
	return (
		<div className='flex justify-center items-center h-screen flex-col'>
			<h3>Welcome, Mr. Guest</h3>
			<div>
				<Link to='/amenities'>Amenities</Link>
			</div>
			<div>
				<Link to='/'>Log out</Link>
			</div>
		</div>
	)
}

export default HomePage