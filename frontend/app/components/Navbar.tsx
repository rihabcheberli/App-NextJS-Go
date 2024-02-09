"use client";

import { useRouter } from 'next/navigation';

const Navbar = () => {
    const router = useRouter();

    const handleLogout = () => {
        router.push('/login');
    };

    return (
        <div className="navbar bg-white shadow-md">
            <div className="flex-1">
                <a className="btn btn-ghost text-xl text-neutral font-bold">Dashboard</a>
            </div>
            <div className="flex-none gap-2">

                <div className="dropdown dropdown-end">
                    <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
                        <div className="w-24 rounded-full ring ring-neutral ring-offset-base-100 ring-offset-2">
                            <img src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                        </div>
                    </div>
                    <ul tabIndex={0} className="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
                        <button className="btn btn-sm btn-ghost" onClick={handleLogout}>  Logout
                        </button>
                    </ul>
                </div>
            </div>
        </div>
    );
};

export default Navbar;
