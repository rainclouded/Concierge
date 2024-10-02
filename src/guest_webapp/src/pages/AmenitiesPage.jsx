import React from 'react'
import { Link } from 'react-router-dom'

const AmenitiesPage = () => {

	return (
		<div className='flex justify-center items-center h-screen flex-col'>
      <h1>Amenities</h1>
			<p>List of amenities</p>
			<Link to='/home'>Homepage</Link>
    </div>
	)
}

export default AmenitiesPage