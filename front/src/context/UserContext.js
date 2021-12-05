import { useState, createContext } from "react";

export const UserContext = createContext();
export const UserProvider = props => {
  const [signed, setSigned] = useState(false);
  return (
    <UserContext.Provider value={[signed, setSigned]}>
      {props.children}
    </UserContext.Provider>
  );
};
