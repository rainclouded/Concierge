import React from 'react'
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {
	const navigate = useNavigate();

	const handleLogin = (e) => {
		e.preventDefault();
		navigate('/home');
		console.log('Logge in')
	}

	return (
		<div className='flex justify-center items-center h-screen flex-col'>
			<h3 className='text-3xl mb-4'>Concierge</h3>
			<form onSubmit={handleLogin} className='border p-5 rounded-xl'>
				<div className='mb-3'>
					<label htmlFor="room-num-input" className='block'>Room Number</label>
					<input id='room-num-input' type="text" placeholder='Your Room Number' className='px-3 py-2 w-full rounded-lg border border-neutral-200'/>
				</div>
				<div className='mb-3'>
					<label htmlFor="pass-code-input" className='block'>Passcode</label>
					<input id='pass-code-input' type="password" placeholder='Your Room Number' className='px-3 py-2 w-full rounded-lg border border-neutral-200'/>
				</div>
				<div className='grid place-content-center'>
					<button type='submit' className='px-3 py-2 bg-neutral-200 rounded-lg'>Sign In</button>
				</div>
			</form>
		</div>
	)
}

export default LoginPage