"use client";

import React, { useState, useEffect } from 'react';
import { USERS_API_URL } from '../api';

const UsersTable = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [updateUserId, setUpdateUserId] = useState(null);
  const [updateEmail, setUpdateEmail] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [newUserEmail, setNewUserEmail] = useState('');
  const [newUserPassword, setNewUserPassword] = useState('');
  const [updateUserName, setUpdateUserName] = useState('');
  const [updateUserLastName, setUpdateUserLastName] = useState('');
  const [newUserName, setNewUserName] = useState('');
  const [newUserLastName, setNewUserLastName] = useState('');


  const fetchUsers = async () => {
    try {
      const response = await fetch(USERS_API_URL);
      if (!response.ok) {
        throw new Error('Failed to fetch users');
      }
      const data = await response.json();
      setUsers(data);
      setLoading(false);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };

  const deleteUser = async (userId) => {
    try {
      const response = await fetch(`${USERS_API_URL}/${userId}`, {
        method: 'DELETE',
      });
      if (!response.ok) {
        throw new Error('Failed to delete user');
      }
      setUsers(users.filter((user) => user.id !== userId));
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  };

  const updateUser = async (userId, newEmail, newName, newLastName) => {
    try {
      const response = await fetch(`${USERS_API_URL}/${userId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: newEmail, name: newName, last_name: newLastName }),
      });
      if (!response.ok) {
        throw new Error('Failed to update user');
      }

      const userIndex = users.findIndex(user => user.id === userId);
      if (userIndex === -1) {
        throw new Error('User not found in the state');
      }

      const updatedUsers = [...users];
      updatedUsers[userIndex] = { ...updatedUsers[userIndex], email: newEmail, name: newName, last_name: newLastName };

      setUsers(updatedUsers);

      setUpdateUserId(null);
      setUpdateEmail('');
    } catch (error) {
      console.error('Error updating user:', error);
    }
  };

  const createUser = async () => {
    try {
      const newUser = { email: newUserEmail, password: newUserPassword, name: newUserName, last_name: newUserLastName };
      const response = await fetch(USERS_API_URL, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newUser),
      });
      if (!response.ok) {
        throw new Error('Failed to create user');
      }
      setIsModalOpen(false);
      setNewUserEmail('');
      setNewUserPassword('');
      setNewUserName('');
      setNewUserLastName('');
      fetchUsers();
    } catch (error) {
      console.error('Error creating user:', error);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  if (loading) {
    return <span className="loading loading-spinner text-neutral"></span>;
  }

  return (
    <div className="flex flex-col items-center justify-center ">
      <div className="m-8">
        <table className="table-auto border-separate border border-slate-400  border-neutral">
          <thead>
            <tr className="bg-neutral">
              <th className="px-4 py-2 text-white">Name</th>
              <th className="px-4 py-2 text-white">Last Name</th>
              <th className="px-4 py-2 text-white">Email</th>
              <th className="px-4 py-2 text-white">Action</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr key={user.id} className="hover:bg-gray-100">
                <td className="px-4 py-2">
                  {updateUserId === user.id ? (
                    <input
                      type="text"
                      value={updateUserName}
                      onChange={(e) => setUpdateUserName(e.target.value)}
                    />
                  ) : (
                    user.name
                  )}
                </td>
                <td className="px-4 py-2">
                  {updateUserId === user.id ? (
                    <input
                      type="text"
                      value={updateUserLastName}
                      onChange={(e) => setUpdateUserLastName(e.target.value)}
                    />
                  ) : (
                    user.last_name
                  )}
                </td>
                <td className="px-4 py-2">
                  {updateUserId === user.id ? (
                    <input
                      type="text"
                      value={updateEmail}
                      onChange={(e) => setUpdateEmail(e.target.value)}
                    />
                  ) : (
                    user.email
                  )}
                </td>
                <td className="px-4 py-2">
                  {updateUserId === user.id ? (
                    <button
                      className="btn btn-sm btn-neutral"
                      onClick={() => updateUser(user.id, updateEmail, updateUserName, updateUserLastName)}
                    >
                      Update
                    </button>
                  ) : (
                    <button
                      className="btn btn-sm btn-neutral mr-2"
                      onClick={() => {
                        setUpdateUserId(user.id);
                        setUpdateEmail(user.email);
                        setUpdateUserName(user.name);
                        setUpdateUserLastName(user.last_name);
                      }}
                    >
                      Edit
                    </button>
                  )}
                  <button
                    className="btn btn-sm btn-neutral"
                    onClick={() => deleteUser(user.id)}
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>

        </table>
        <button
          className="btn btn-primary mt-4"
          onClick={() => setIsModalOpen(true)}
        >
          Create User
        </button>
      </div>
      {isModalOpen && (
        <div className="fixed top-0 left-0 w-full h-full flex justify-center items-center bg-gray-800 bg-opacity-50">
          <div className="bg-white p-4 rounded shadow-md ">
            <h2 className="text-lg font-semibold mb-4 ">Create User</h2>
            <input
              type="text"
              placeholder="Name"
              value={newUserName}
              onChange={(e) => setNewUserName(e.target.value)}
              className="input w-full mb-2"
            />
            <input
              type="text"
              placeholder="Last Name"
              value={newUserLastName}
              onChange={(e) => setNewUserLastName(e.target.value)}
              className="input w-full mb-2"
            />

            <input
              type="text"
              placeholder="Email"
              value={newUserEmail}
              onChange={(e) => setNewUserEmail(e.target.value)}
              className="input w-full mb-2"
            />
            <input
              type="password"
              placeholder="Password"
              value={newUserPassword}
              onChange={(e) => setNewUserPassword(e.target.value)}
              className="input w-full mb-2"
            />
            <div className="flex justify-end">
              <button
                className="btn btn-primary mr-2"
                onClick={createUser}
              >
                Create
              </button>
              <button
                className="btn btn-neutral"
                onClick={() => setIsModalOpen(false)}
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default UsersTable;
