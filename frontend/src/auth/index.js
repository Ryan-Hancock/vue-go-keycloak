import axios from "axios"

export default {

    client: axios.create({
        baseURL: 'http://localhost:8090',
        timeout: 1000,
      }),

    getToken() {
        if (localStorage.getItem('token') == null) {
            return false
        }

        return JSON.parse(localStorage.getItem('token'))
    }, 
  
    login({username, password}) {
        return this.client.post('/auth/login', {
                username,
                password
        }).then(response => {
            localStorage.setItem('token', JSON.stringify(response.data))
            return response
        }).catch(error => {
            return error
        })
    },
  
    signup({email, username, password}) {
        return this.client.post('/auth/create', {
            email,
            username,
            password
        }).then(response => {
            return response.data
        }).catch(error => {
            return error
        })
    },
  
    logout() {
      localStorage.removeItem('token')
    },
  
    checkAuth() {
        if (!this.getToken()) {
            return Promise.reject()
        }
        
        return this.client.post('/auth/validate', {
            accessToken: this.getToken().access_token
        }).then(response => {
            return response.data
        }).catch(error => {
            return Promise.reject()
        })
    },


    refresh() {
        if (!this.getToken()) {
            return Promise.reject()
        }

        return this.client.post('/auth/refresh', {
            refreshToken: this.getToken().refresh_token
        }).then(response => {
            localStorage.setItem('token', JSON.stringify(response.data))
            return response
        }).catch(error => {
            return Promise.reject()
        })
    },
  
  
    //future use with another service
    getAuthHeader() {
      return {
        'Authorization': 'Bearer ' + localStorage.getItem('token')
      }
    }
  }