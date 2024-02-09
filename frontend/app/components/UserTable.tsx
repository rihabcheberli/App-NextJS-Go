"use client";

import React, { useState, useEffect } from 'react';
import { USERS_API_URL } from '../api'; 

const UsersTable = () => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);

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

  useEffect(() => {
    fetchUsers();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h2>Users Table</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Email</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id}>
              <td>{user.id}</td>
              <td>{user.email}</td>
              <td>
                <button onClick={() => deleteUser(user.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default UsersTable;
