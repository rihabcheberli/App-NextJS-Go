import React from 'react';
import Login from '../components/Login';
import Link from 'next/link';

const LoginPage = () => {
    return (<div>
        <Login />
        <div className="text-sm text-center">
            Don't have an account?{' '}
            <Link href="/register">
                <div className="text-info">Register</div>
            </Link>
        </div></div>

    );
};

export default LoginPage;
