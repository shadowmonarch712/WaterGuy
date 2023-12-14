<a name="readme-top"></a>
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="https://i.imgur.com/tiRmIty.png" alt="Logo" width="650" height="320">
  </a>

<h1 align="center">WaterGuy</h1>

  <p align="center">
    <h2>With WaterGuy Never Forget to Hydrate yourself while solving those bugs</h2>
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Welcome to Water Guy, a unique and innovative application designed to help you stay hydrated throughout the day. Built with a robust GoFr backend and a dynamic React.js frontend, Water Guy is more than just a water reminder app - it’s a comprehensive hydration management system.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![React][React.js]][React-url]
* [![goFr][goFr]][goFr-url]
* [![MongoDB][MongoDB]][MongoDB-url]
* [![TailwindCSS][TailwindCSS]][TailwindCSS-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* npm
  ```sh
  npm install npm@latest -g
  ```
* go

### Installation


1. Clone the repo
   ```sh
   git clone https://github.com/shadowmonarch712/waterguy.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```
3. Run frontend
   ```sh
   npm run dev
   ```   
5. Enter your MongoDB credentials in `config.js`
   ```js
   MONGO_DB_HOST= 'ENTER YOUR HOST' 
   MONGO_DB_PORT= 'ENTER MONGODB PORT'
   MONGO_DB_USER= 'ENTER USERNAME'
   MONGO_DB_PASS= 'ENTER PASSWORD'
   MONGO_DB_NAME= 'ENTER DATABASE NAME'
   ```
6. Enter MongoDB URI in `database/db.go`
   ```go
   mongoURI := "mongodb://username:password@host:port/?authMechanism"
   ```
7. Download necessary modules
   ```go
   go mod tidy
   ```
8. Run server
   ```go
   go run main.go
   ```  
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

WaterGuy can be used to create a goal : 
<div align="center"> 
  <img src="https://i.imgur.com/PDgmNJf.png" alt="Logo" width="320" height="320" style="display: inline-block;">
  <img src="https://i.imgur.com/zedslJX.png" alt="Logo" width="500" height="320" style="display: inline-block;">
</div>

WaterGuy can be used to fetch goals : 
<div align="center"> 
  <img src="https://i.imgur.com/K3SJ5Ta.png" alt="Logo" width="320" height="320" style="display: inline-block;">
  <img src="https://i.imgur.com/IhjfQLw.png" alt="Logo" width="500" height="320" style="display: inline-block;">
</div>

WaterGuy can be used to update or delete goals : 
<div align="center"> 
  <img src="https://i.imgur.com/rOHcM3F.png" alt="Logo" width="320" height="320" ">
  <img src="https://i.imgur.com/o6UbxR9.png" alt="Logo" width="500" height="320" style="display: inline-block;">
  <img src="https://i.imgur.com/hUVWYTN.png" alt="Logo" width="320" height="320" style="display: inline-block;">
</div>


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] CRUD Functionality <br>
      At the core of Water Guy is its CRUD functionality. Users can create their own hydration goals, read or fetch any previous goals, update them as needed, and delete them when they’re no longer relevant.           This gives users full control over their hydration plan, allowing for a truly personalized experience.
- [ ] Water Reminder <br>
      Once a goal is set, users can start their daily water reminder. This feature asks for an ‘hour multiplier’ - a user-defined interval that determines how often they should drink a glass of water. After each       interval, the goal count is decremented, providing a clear and simple way to track water intake.
- [ ] Audio Alerts <br>
      To ensure users never miss a hydration break, Water Guy plays an audio alert after every interval. This friendly reminder serves as a nudge to drink water, helping users meet their hydration goals with           ease.

Water Guy is more than an app - it’s a lifestyle choice. By promoting regular hydration, it supports overall health and wellbeing. So why wait? Start your hydration journey with Water Guy today!

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Aditya Srivastava - adityasrivastava0709@gmail.com

Project Link: [https://github.com/shadowmonarch712/waterguy](https://github.com/shadowmonarch712/waterguy)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[React.js]: https://i.imgur.com/oTR5pw8.png
[React-url]: https://reactjs.org/
[goFr]: https://i.imgur.com/ecN4cxl.png
[goFr-url]: https://gofr.dev/
[MongoDB]: https://i.imgur.com/eV7Ee9T.png
[MongoDB-url]: https://www.mongodb.com/
[TailwindCSS]: https://i.imgur.com/3p7TTje.png
[TailwindCSS-url]: https://tailwindcss.com/

