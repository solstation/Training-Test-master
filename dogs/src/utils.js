const dogsName = 'local-storage-dogs'

// Return true if a url is a video
const isVideo = url => {
    let pattern = new RegExp('https?.*?\.mp4')

    return pattern.test(url)
}

// Return true if the dogs object exists in local storage
const existsDogs = ()=> {
    return localStorage.getItem(dogsName) !== null
}

// Return the dogs
const getDogs = ()=>{
    return existsDogs() ? JSON.parse(localStorage.getItem(dogsName)) : []
}

// Save a dog in local storage
const saveDog = dog => {
    if (!existsDogs()) localStorage.setItem(dogsName, JSON.stringify([]))

    let dogs = getDogs()
    dogs.push(dog)

    localStorage.setItem(dogsName, JSON.stringify(dogs))
}

// Remove all dogs from local storage
const clearDogs = ()=>{
    localStorage.removeItem(dogsName)
}

export default {
    isVideo, saveDog, clearDogs, getDogs
}