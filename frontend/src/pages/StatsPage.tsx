import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { 
  Box, 
  Card, 
  CardBody, 
  CardHeader, 
  Heading, 
  Spinner, 
  Text, 
  useToast 
} from '@chakra-ui/react';
import { URLStats } from '../types/url';
import { urlAPI } from '../services/api';
import { handleApiError } from '../utils/helpers';
import URLCard from '../components/URLCard';

const StatsPage = () => {
  const { shortCode } = useParams<{ shortCode: string }>();
  const toast = useToast();
  
  const [urlStats, setUrlStats] = useState<URLStats | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchUrlStats = async () => {
      if (!shortCode) return;

      setIsLoading(true);
      try {
        const result = await urlAPI.getURLStats(shortCode);
        setUrlStats(result);
      } catch (error) {
        const errorMessage = handleApiError(error);
        setError(errorMessage);
        toast({
          title: 'Error loading URL statistics',
          description: errorMessage,
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
      } finally {
        setIsLoading(false);
      }
    };

    fetchUrlStats();
  }, [shortCode, toast]);

  if (isLoading) {
    return (
      <Box textAlign="center" py={10}>
        <Spinner size="xl" />
        <Text mt={4}>Loading URL statistics...</Text>
      </Box>
    );
  }

  if (error || !urlStats) {
    return (
      <Box textAlign="center" py={10}>
        <Heading size="md" color="red.500">Error Loading Statistics</Heading>
        <Text mt={4}>{error || 'URL not found'}</Text>
      </Box>
    );
  }

  return (
    <Box>
      <Card mb={6}>
        <CardHeader>
          <Heading size="md">URL Statistics</Heading>
          <Text mt={2} color="gray.600">
            Detailed information for short code: <strong>{shortCode}</strong>
          </Text>
        </CardHeader>
        <CardBody>
          <URLCard urlData={urlStats} showStats={true} />
        </CardBody>
      </Card>
    </Box>
  );
};

export default StatsPage; 