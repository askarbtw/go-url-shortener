import { Box, Container, Flex, Heading, Text } from '@chakra-ui/react';
import { Link as RouterLink, Outlet } from 'react-router-dom';

const Layout = () => {
  return (
    <Box minH="100vh" bg="gray.50">
      <Flex 
        as="header" 
        bg="blue.600" 
        color="white" 
        p={4} 
        shadow="md"
      >
        <Container maxW="container.xl">
          <Flex justifyContent="space-between" alignItems="center">
            <RouterLink to="/" style={{ textDecoration: 'none' }}>
              <Heading size="lg" color="white">URL Shortener</Heading>
            </RouterLink>
          </Flex>
        </Container>
      </Flex>

      <Container maxW="container.xl" py={8}>
        <Outlet />
      </Container>

      <Box as="footer" bg="gray.100" p={4} mt="auto">
        <Container maxW="container.xl">
          <Flex justifyContent="center">
            <Text fontSize="sm" color="gray.600">
              &copy; {new Date().getFullYear()} URL Shortener Service
            </Text>
          </Flex>
        </Container>
      </Box>
    </Box>
  );
};

export default Layout; 