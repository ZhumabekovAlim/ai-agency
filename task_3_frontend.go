```jsx
import React, { useState } from 'react';

const AdminPanel = () => {
  const [orders, setOrders] = useState([]);
  const [inventory, setInventory] = useState([]);
  const [newOrder, setNewOrder] = useState('');
  const [newInventoryItem, setNewInventoryItem] = useState('');

  const handleOrderChange = (e) => setNewOrder(e.target.value);
  const handleInventoryChange = (e) => setNewInventoryItem(e.target.value);

  const addOrder = () => {
    setOrders([...orders, newOrder]);
    setNewOrder('');
  };

  const addInventory = () => {
    setInventory([...inventory, newInventoryItem]);
    setNewInventoryItem('');
  };

  return (
    <div>
      <h1>Admin Panel</h1>
      <div>
        <h2>Orders</h2>
        <input
          placeholder="New Order"
          value={newOrder}
          onChange={handleOrderChange}
        />
        <button onClick={addOrder}>Add Order</button>
        <ul>
          {orders.map((order, index) => (
            <li key={index}>{order}</li>
          ))}
        </ul>
      </div>
      <div>
        <h2>Inventory</h2>
        <input
          placeholder="New Inventory Item"
          value={newInventoryItem}
          onChange={handleInventoryChange}
        />
        <button onClick={addInventory}>Add Inventory</button>
        <ul>
          {inventory.map((item, index) => (
            <li key={index}>{item}</li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default AdminPanel;
```