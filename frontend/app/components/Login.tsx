"use client";

import React, { useState } from 'react';
import { LOGIN_API_URL } from '../api';
import { useRouter } from 'next/navigation'
import Link from 'next/link';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const router = useRouter()

  const handleLogin = async () => {
    try {
      const response = await fetch(LOGIN_API_URL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });
      const data = await response.json();
      if (response.ok) {
        router.push('/dashboard');
      } else {
        setError(data.message);
      }
    } catch (error) {
      setError('Failed to login. Please try again.');
    }
  };

  return (
    <div className="flex justify-center items-center p-8">
      <div className="p-8 bg-white  shadow-2xl rounded-md w-3/4">
        <div className="text-2xl text-neutral text-center mb-4">Login</div>
        {error && <div className="text-error mb-4">{error}</div>}
        <input
          className="input input-bordered w-full mb-8 mt-2"
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          className="input input-bordered w-full"
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="flex flex-col items-center">
          <button
            className="btn btn-active btn-neutral text-white w-1/2 max-w-xs m-8"
            onClick={handleLogin}
          >
            Login
          </button>
        </div>
        <div className="text-sm text-center">
          Don't have an account?{' '}
          <Link href="/register">
            <div className="text-info">Register</div>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Login;
