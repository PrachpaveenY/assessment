import React from 'react'
import './App.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.bundle.min.js'

function App() {

  return (
    <div className="App">
      <nav>
        <a className='navbar-expand-sm dark'>
          <img src='./assets/reactjs.png' height='30' alt=''/>
        </a>
        <span className='navbar-brand text-info'>React.js</span>
        <ul>
          <li className='nav-item'>
            <a className='nav-link'>Home</a>
          </li>
          <li className='nav-item'>
            <a className='nav-link'>Product</a>
          </li>
          <li className='nav-item'>
            <a className='nav-link'>Member</a>
          </li>
          <li className='nav-item dropdown'>
            <a className='nav-link dropdown' data-bs-toggle='dropdown'>javascript</a>
            <div className='dropdown menu'>
              <a className='dropdown-item'>React</a>
              <a className='dropdown-item'>React</a>
              <a className='dropdown-item'>React</a>
            </div>
          </li>
        </ul>
      </nav>
      <table border="1">
      <tr><th>Product</th><th>Price</th></tr>
      <tr><th>Product</th><th>Price</th></tr>
      <tr><th>Product</th><th>Price</th></tr>
      </table>
      <div>
        <div className='text-primary bg-warning'>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Maxime fugit molestiae ad, eveniet esse dolore ratione doloribus delectus, ipsum quos sequi voluptatem iste quisquam impedit. Cum voluptas ea quisquam nobis!</div>
      </div>
    </div>
  )
}

export default App
