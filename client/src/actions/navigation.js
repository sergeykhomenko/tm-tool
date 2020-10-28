export const navigate = (page, options) => ({type: 'NAVIGATE_TO', payload: {page, options}})

export const openProject = projectId => ({type: 'OPEN_PROJECT', payload: projectId})

export const back = () => ({type: 'NAVIGATE_BACK'})