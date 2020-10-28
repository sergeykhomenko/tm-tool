function get(type, url) {
    return {
        type,
        payload: {
            request: { url }
        }
    }
}

function post(type, url, data) {
    return {
        type,
        payload: {
            request: {
                method: 'POST',
                url,
                data
            }
        }
    }
}

export const listProjects = () => get('LIST_PROJECTS', '/projects')
export const createProject = (title, description) => post('CREATE_PROJECT', '/projects', {title, description})
export const listTickets = (title, description) => true
export const createTicket = (projectId, title, description) => post('CREATE_TICKET', `/projects/${projectId}`, {title, description})
export const updateTicket = (title, description) => true

export const listMessage = (title, description) => true
export const createMessage = (title, description) => true