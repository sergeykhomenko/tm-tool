import React from 'react';
import ProjectContainer from "./components/containers/ProjectContainer";
import TicketContainer from "./components/containers/TIcketContainer";
import TicketViewContainer from "./components/containers/TicketViewContainer";
import {useSelector} from "react-redux";
import {Title} from "./components/headings";
import {ContentBox} from "./components/wrappers";
import {Button} from "./components/buttons";

const AppLegacy = () => {
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

const App = () => {
  return <ContentBox>
    <div className='header'>
      <Title label='Time Management Tool' />
      <div className="header__actions">
        <Button />
      </div>
    </div>
    
    <div className="list">
      <div className="list-item"></div>
      <div className="list-item"></div>
      <div className="list-item"></div>
    </div>
  </ContentBox>;
}

export default App;
