import { useEffect, useState, forwardRef, useImperativeHandle } from 'react';
import { 
  Box, 
  Heading, 
  Spinner, 
  Text, 
  SimpleGrid, 
  useToast, 
  Button,
  Icon,
  Flex,
  Badge
} from '@chakra-ui/react';
import { FiBarChart2 } from 'react-icons/fi';
import { useNavigate } from 'react-router-dom';
import { URLStats } from '../types/url';
import { urlAPI } from '../services/api';
import { handleApiError } from '../utils/helpers';
import URLCard from './URLCard';

// Define the ref handle type
export interface URLListRefHandle {
  fetchUrls: () => Promise<void>;
}

const URLList = forwardRef<URLListRefHandle, {}>((props, ref) => {
  const [urls, setUrls] = useState<URLStats[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');
  const toast = useToast();
  const navigate = useNavigate();

  const fetchUrls = async () => {
    setIsLoading(true);
    try {
      const urlsData = await urlAPI.getAllURLs();
      setUrls(urlsData);
    } catch (error) {
      const errorMessage = handleApiError(error);
      setError(errorMessage);
      toast({
        title: 'Error loading URLs',
        description: errorMessage,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    } finally {
      setIsLoading(false);
    }
  };

  // Expose the fetchUrls method to parent components
  useImperativeHandle(ref, () => ({
    fetchUrls
  }));

  useEffect(() => {
    fetchUrls();
  }, []);

  const handleDelete = async (shortCode: string) => {
    try {
      await urlAPI.deleteURL(shortCode);
      // Remove the deleted URL from the state
      setUrls(urls.filter(url => url.shortCode !== shortCode));
      return true;
    } catch (error) {
      const errorMessage = handleApiError(error);
      toast({
        title: 'Error deleting URL',
        description: errorMessage,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
      return false;
    }
  };

  const goToStats = (shortCode: string) => {
    navigate(`/stats/${shortCode}`);
  };

  if (isLoading) {
    return (
      <Box textAlign="center" py={10}>
        <Spinner size="xl" />
        <Text mt={4}>Loading URLs...</Text>
      </Box>
    );
  }

  if (error) {
    return (
      <Box textAlign="center" py={10}>
        <Heading size="md" color="red.500">Error Loading URLs</Heading>
        <Text mt={4}>{error}</Text>
        <Button mt={4} onClick={fetchUrls}>Try Again</Button>
      </Box>
    );
  }

  if (urls.length === 0) {
    return (
      <Box textAlign="center" py={10} bg="gray.50" borderRadius="md" p={6}>
        <Heading size="md">No URLs Found</Heading>
        <Text mt={4}>Create your first short URL above.</Text>
      </Box>
    );
  }

  return (
    <Box>
      <SimpleGrid columns={{ base: 1, md: 2 }} spacing={6}>
        {urls.map(url => (
          <Box key={url.shortCode} position="relative">
            <URLCard urlData={url} onDelete={handleDelete} />
            <Flex 
              position="absolute" 
              top="0.5rem" 
              right="0.5rem"
              alignItems="center"
            >
              <Badge 
                colorScheme="blue" 
                display="flex" 
                alignItems="center" 
                p={1} 
                borderRadius="md"
                cursor="pointer"
                onClick={() => goToStats(url.shortCode)}
              >
                <Icon as={FiBarChart2} mr={1} />
                {url.accessCount} clicks
              </Badge>
            </Flex>
          </Box>
        ))}
      </SimpleGrid>
    </Box>
  );
});

URLList.displayName = 'URLList';

export default URLList; 