import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
  const token = request.cookies.get('token');
  const { pathname } = request.nextUrl;

  // Define protected routes
  const protectedRoutes = ['/search'];
  const authRoutes = ['/login', '/register'];

  // Check if the route is protected and user is not authenticated
  if (protectedRoutes.includes(pathname) && !token) {
    return NextResponse.redirect(new URL('/login', request.url));
  }

  // Redirect authenticated users away from auth pages
  if (authRoutes.includes(pathname) && token) {
    return NextResponse.redirect(new URL('/search', request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ['/search', '/login', '/register'],
};
