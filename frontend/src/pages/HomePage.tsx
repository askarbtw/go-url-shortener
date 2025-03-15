import { useState } from 'react';
import { 
  Box, 
  Heading, 
  Text, 
  useToast,
  VStack,
  Card,
  CardHeader,
  CardBody
} from '@chakra-ui/react';
import { urlAPI } from '../services/api';
import { handleApiError } from '../utils/helpers';
import { URL } from '../types/url';
import URLForm from '../components/URLForm';
import URLCard from '../components/URLCard';

const HomePage = () => {
  const toast = useToast();
  const [isLoading, setIsLoading] = useState(false);
  const [shortenedUrl, setShortenedUrl] = useState<URL | null>(null);

  const handleCreateURL = async (url: string) => {
    setIsLoading(true);
    try {
      const result = await urlAPI.createURL({ url });
      setShortenedUrl(result);
      toast({
        title: 'URL shortened successfully',
        status: 'success',
        duration: 3000,
        isClosable: true,
      });
    } catch (error) {
      const errorMessage = handleApiError(error);
      toast({
        title: 'Error shortening URL',
        description: errorMessage,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Box>
      <Card mb={6}>
        <CardHeader>
          <Heading size="md">Shorten a URL</Heading>
          <Text mt={2} color="gray.600">
            Enter a long URL to create a shorter, more manageable link.
          </Text>
        </CardHeader>
        <CardBody>
          <URLForm
            onSubmit={handleCreateURL}
            isLoading={isLoading}
            buttonText="Shorten URL"
          />
        </CardBody>
      </Card>

      {shortenedUrl && (
        <VStack spacing={4} align="stretch">
          <Heading size="md">Your Shortened URL</Heading>
          <URLCard urlData={shortenedUrl} />
        </VStack>
      )}
    </Box>
  );
};

export default HomePage; 