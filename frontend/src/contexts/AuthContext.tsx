import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import axios from 'axios';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';

interface AuthContextType {
  user: any;
  login: (phoneNumber: string, password: string) => Promise<void>;
  register: (name: string, phoneNumber: string, password: string) => Promise<void>;
  logout: () => void;
  isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<any>(null);
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    // Check if user is logged in
    const token = Cookies.get('token');
    if (token) {
      setUser({ token });
    }
    setIsLoading(false);
  }, []);

  const login = async (phoneNumber: string, password: string) => {
    try {
      const response = await axios.post('http://localhost:8080/api/login', {
        phone_number: phoneNumber,
        password,
      });
      const { token } = response.data;
      Cookies.set('token', token, { expires: 7 }); // Expires in 7 days
      setUser({ token });
      router.push('/search');
    } catch (error) {
      throw error;
    }
  };

  const register = async (name: string, phoneNumber: string, password: string) => {
    try {
      await axios.post('http://localhost:8080/api/register', {
        name,
        phone_number: phoneNumber,
        password,
      });
      // After successful registration, log the user in
      await login(phoneNumber, password);
    } catch (error) {
      throw error;
    }
  };

  const logout = () => {
    Cookies.remove('token');
    setUser(null);
    router.push('/');
  };

  return (
    <AuthContext.Provider value={{ user, login, register, logout, isLoading }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
