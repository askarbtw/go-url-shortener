import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
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
import { URL } from '../types/url';
import { urlAPI } from '../services/api';
import { handleApiError } from '../utils/helpers';
import URLForm from '../components/URLForm';

const EditPage = () => {
  const { shortCode } = useParams<{ shortCode: string }>();
  const navigate = useNavigate();
  const toast = useToast();
  
  const [url, setUrl] = useState<URL | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [isUpdating, setIsUpdating] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchUrl = async () => {
      if (!shortCode) return;

      setIsLoading(true);
      try {
        const result = await urlAPI.getURL(shortCode);
        setUrl(result);
      } catch (error) {
        const errorMessage = handleApiError(error);
        setError(errorMessage);
        toast({
          title: 'Error loading URL',
          description: errorMessage,
          status: 'error',
          duration: 5000,
          isClosable: true,
        });
      } finally {
        setIsLoading(false);
      }
    };

    fetchUrl();
  }, [shortCode, toast]);

  const handleUpdateURL = async (newUrl: string) => {
    if (!shortCode) return;

    setIsUpdating(true);
    try {
      await urlAPI.updateURL(shortCode, { url: newUrl });
      toast({
        title: 'URL updated successfully',
        status: 'success',
        duration: 3000,
        isClosable: true,
      });
      navigate('/');
    } catch (error) {
      const errorMessage = handleApiError(error);
      toast({
        title: 'Error updating URL',
        description: errorMessage,
        status: 'error',
        duration: 5000,
        isClosable: true,
      });
    } finally {
      setIsUpdating(false);
    }
  };

  if (isLoading) {
    return (
      <Box textAlign="center" py={10}>
        <Spinner size="xl" />
        <Text mt={4}>Loading URL information...</Text>
      </Box>
    );
  }

  if (error || !url) {
    return (
      <Box textAlign="center" py={10}>
        <Heading size="md" color="red.500">Error Loading URL</Heading>
        <Text mt={4}>{error || 'URL not found'}</Text>
      </Box>
    );
  }

  return (
    <Box>
      <Card mb={6}>
        <CardHeader>
          <Heading size="md">Edit URL</Heading>
          <Text mt={2} color="gray.600">
            Update the destination for short code: <strong>{shortCode}</strong>
          </Text>
        </CardHeader>
        <CardBody>
          <URLForm
            initialUrl={url.url}
            onSubmit={handleUpdateURL}
            isLoading={isUpdating}
            buttonText="Update URL"
          />
        </CardBody>
      </Card>
    </Box>
  );
};

export default EditPage; 