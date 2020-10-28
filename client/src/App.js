import React from 'react';
import ProjectContainer from "./components/containers/ProjectContainer";
import TicketContainer from "./components/containers/TIcketContainer";
import TicketViewContainer from "./components/containers/TicketViewContainer";
import {useSelector} from "react-redux";

const App = () => {
  const breadcrumbs = useSelector(state => state.breadcrumbs);

  if (!Array.isArray(breadcrumbs) || breadcrumbs.length === 0) {
    return <ProjectContainer />;
  }

  switch (breadcrumbs[breadcrumbs.length - 1]) {
    case 'project':
      return <TicketContainer />;
    case 'ticket':
      return <TicketViewContainer />;
    default:
      return <ProjectContainer />;
  }
}

export default App;
