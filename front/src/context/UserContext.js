import { useState, createContext } from "react";

export const UserContext = createContext();

export const UserProvider = props => {
  const [user, setUser] = useState(undefined);
  return (
    <UserContext.Provider value={[user, setUser]}>
      {props.children}
    </UserContext.Provider>
  );
};
