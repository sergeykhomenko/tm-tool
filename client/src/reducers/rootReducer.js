export const initialState = {
    breadcrumbs: [],
    currentProject: 0,
    projects: []
};

export const rootReducer = (state = initialState, action) => {
    switch (action.type) {
        case 'LIST_PROJECTS_SUCCESS':
            return { ...state, projects: action.payload.data.projects }
        case 'CREATE_PROJECT_SUCCESS':
            state.projects.push(action.payload.data.project);
            return { ...state, projects: [...state.projects]}
        case 'OPEN_PROJECT':
            state.breadcrumbs.push('project');
            return { ...state, breadcrumbs: [...state.breadcrumbs], currentProject: action.payload }
        case 'NAVIGATE_BACK':
            state.breadcrumbs.pop();
            return { ...state, breadcrumbs: [...state.breadcrumbs], currentProject: action.payload }
        default:
            return { ...state }
    }
}