import axios from 'axios'
import configService from './config'

const dogsService = {}

// Fetch the random dog image
dogsService.getDog = function(){
    return axios.get(configService.apiURL, {})
    .then(res => res.data)
}

export default dogsService