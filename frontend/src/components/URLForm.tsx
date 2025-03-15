import { useState } from 'react';
import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  VStack,
} from '@chakra-ui/react';

interface URLFormProps {
  initialUrl?: string;
  onSubmit: (url: string) => void;
  isLoading: boolean;
  buttonText: string;
}

const URLForm = ({ initialUrl = '', onSubmit, isLoading, buttonText }: URLFormProps) => {
  const [url, setUrl] = useState(initialUrl);
  const [error, setError] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    // Basic validation
    if (!url) {
      setError('URL is required');
      return;
    }

    // Reset error
    setError('');
    
    // Submit the URL
    onSubmit(url);
  };

  return (
    <form onSubmit={handleSubmit}>
      <VStack spacing={4} align="stretch">
        <FormControl isInvalid={!!error}>
          <FormLabel htmlFor="url">Enter URL</FormLabel>
          <Input
            id="url"
            type="text"
            placeholder="https://example.com"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            isDisabled={isLoading}
          />
          {error && <FormErrorMessage>{error}</FormErrorMessage>}
        </FormControl>

        <Button
          type="submit"
          colorScheme="blue"
          isLoading={isLoading}
          isDisabled={isLoading}
          width="full"
        >
          {buttonText}
        </Button>
      </VStack>
    </form>
  );
};

export default URLForm; 