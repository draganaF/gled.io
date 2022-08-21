import axios from "axios";
import { USER_SERVICE_URL } from "../../url";

const state = {
  user: null,
  users: null,
  result: null
};

const getters = {
  getUser: state => state.user,
  getUsers: state => state.users,
  getResult: state => state.result
};

const actions = {

  searchUsers: (context, params) => {
    axios.post(`${USER_SERVICE_URL}/users/search`, params)
      .then(response => {
        console.log(params)
        console.log(response.data);
        context.commit('setUsers', response.data);
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'fetch',
          ok: false,
          message: error.response.data
        });
      })
  },

  fetchUser: (context, userId) => {
    axios.get(`${USER_SERVICE_URL}/users/${userId}`)
      .then(response => {
        console.log(response.data);
        context.commit('setUser', response.data);
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'fetch',
          ok: false,
          message: error.response.data
        });
      })
  },

  activateUser: (context, link) => {
    axios.get(`${USER_SERVICE_URL}/users/activate/${link}`)
      .then(response => {
        context.commit('setResult', {
          label: 'activate',
          ok: true,
          message: 'You have successfully activated the profile.'
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'activate',
          ok: false,
          message: error.response.data
        });
      })
  },

  createUser: (context, user) => {
    axios.post(`${USER_SERVICE_URL}/users`, user)
      .then(response => {
        console.log(response);
        context.commit('setUser', response.data);
        context.commit('setResult', {
          label: 'create',
          ok: true,
          message: 'You have successfully create a new user.'
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'create',
          ok: false,
          message: error.response.data
        });
      });
  },

  updateUser: (context, user) => {
    axios.put(`${USER_SERVICE_URL}/users/update`, user)
      .then(response => {
        console.log(response);
        context.commit('setUser', response.data);
        context.commit('setResult', {
          label: 'update',
          ok: true,
          message: 'You have successfully updated personal informations.'
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'update',
          ok: false,
          message: error.response.data
        });
      });
  },
  updateBalance: (context, balanceObject) => {
    axios.post(`${USER_SERVICE_URL}/users/update-balance`, balanceObject)
      .then(response => {
        console.log(response);
        context.commit('setResult', {
          label: 'balance',
          ok: true,
          message: 'You have successfully updated balanse.'
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'balance',
          ok: false,
          message: error.response.data
        });
      });
  },

  updateUserPassword: (context, credentials) => {
    axios.put(`${USER_SERVICE_URL}/users/update/password`, credentials)
      .then(response => {
        console.log(response);
        context.commit('setUser', response.data);
        context.commit('setResult', {
          label: 'updatePassword',
          ok: true,
          message: 'You have successfully updated credentials.'
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'updatePassword',
          ok: false,
          message: error.response.data
        });
      });
  },

  deleteUser: (context, userId) => {
    axios.delete(`${USER_SERVICE_URL}/users/delete-user/${userId}`)
      .then(response => {
        console.log(response.data);
        context.commit('setResult', {
          label: 'delete',
          ok: true,
          message: `You have successfully deleted a user.`
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'delete',
          ok: false,
          message: error.response.data
        });
      })
  },

  blockUser: (context, userId) => {
    axios.get(`${USER_SERVICE_URL}/users/block-user/${userId}`)
      .then(response => {
        console.log(response.data);
        context.commit('setResult', {
          label: 'block',
          ok: true,
          message: `You have successfully blocked a user.`
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'block',
          ok: false,
          message: error.response.data
        });
      })
  },
  incrementWorkersSoldTickets: (context, userId) => {
    axios.get(`${USER_SERVICE_URL}/users/increment-sold-tickets/${userId}`)
      .then(response => {
        console.log(response.data);
        context.commit('setResult', {
          label: 'increment',
          ok: true,
          message: `You have successfully sold ticket.`
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'increment',
          ok: false,
          message: error.response.data
        });
      })
  },
  incrementUsersBoughtTickets: (context, userId) => {
    axios.get(`${USER_SERVICE_URL}/users/increment-bought-tickets/${userId}`)
      .then(response => {
        context.commit('setResult', {
          label: 'increment',
          ok: true,
          message: `You have successfully sold ticket.`
        });
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'increment',
          ok: false,
          message: error.response.data
        });
      })
  },

}


const mutations = {

  setUsers: (state, response) => {
    state.users = response;
  },

  setUser: (state, response) => {
    state.user = response;
  },

  setResult: (state, response) => {
    state.result = response;
  }
};

export default {
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations,
  namespaced: true
};