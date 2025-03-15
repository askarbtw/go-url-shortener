import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { 
  Box, 
  Card, 
  CardBody, 
  CardHeader, 
  Heading, 
  Spinner, 
  Text, 
  useToast,
  Stat,
  StatLabel,
  StatNumber,
  StatHelpText,
  StatGroup,
  Button,
  Icon,
  Flex,
  Divider,
  Badge,
  SimpleGrid,
  useClipboard
} from '@chakra-ui/react';
import { FiArrowLeft, FiCopy, FiExternalLink } from 'react-icons/fi';
import { URLStats } from '../types/url';
import { urlAPI, getRedirectURL } from '../services/api';
import { handleApiError, formatDate } from '../utils/helpers';
import URLCard from '../components/URLCard';

const StatsPage = () => {
  const { shortCode } = useParams<{ shortCode: string }>();
  const toast = useToast();
  
  const [urlStats, setUrlStats] = useState<URLStats | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');
  
  const redirectUrl = shortCode ? getRedirectURL(shortCode) : '';
  const { hasCopied, onCopy } = useClipboard(redirectUrl);

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
        <Button as={Link} to="/" leftIcon={<FiArrowLeft />} mt={6}>
          Back to Home
        </Button>
      </Box>
    );
  }

  return (
    <Box>
      <Button as={Link} to="/" leftIcon={<FiArrowLeft />} mb={6} size="sm" variant="outline">
        Back to Home
      </Button>

      <Card mb={6}>
        <CardHeader>
          <Heading size="md">URL Statistics</Heading>
          <Text mt={2} color="gray.600">
            Detailed information for short code: <strong>{shortCode}</strong>
          </Text>
        </CardHeader>
        <CardBody>
          <Box mb={6}>
            <Heading size="sm" mb={2}>Original URL</Heading>
            <Text 
              p={2} 
              bg="gray.50" 
              borderRadius="md" 
              fontSize="sm" 
              wordBreak="break-all"
            >
              {urlStats.url}
            </Text>
          </Box>

          <Box mb={6}>
            <Heading size="sm" mb={2}>Short URL</Heading>
            <Flex 
              p={2} 
              bg="blue.50" 
              borderRadius="md" 
              alignItems="center" 
              justifyContent="space-between"
            >
              <Text color="blue.600" fontSize="sm">
                {redirectUrl}
              </Text>
              <Flex>
                <Button 
                  size="xs" 
                  leftIcon={<FiCopy />} 
                  onClick={onCopy} 
                  ml={2}
                >
                  {hasCopied ? "Copied!" : "Copy"}
                </Button>
                <Button 
                  as="a" 
                  href={redirectUrl} 
                  target="_blank" 
                  size="xs" 
                  leftIcon={<FiExternalLink />} 
                  ml={2}
                >
                  Open
                </Button>
              </Flex>
            </Flex>
          </Box>

          <Divider my={6} />

          <SimpleGrid columns={{ base: 1, md: 3 }} spacing={4} mb={6}>
            <Card>
              <CardBody>
                <Stat>
                  <StatLabel>Total Clicks</StatLabel>
                  <StatNumber>{urlStats.accessCount}</StatNumber>
                  <StatHelpText>Since creation</StatHelpText>
                </Stat>
              </CardBody>
            </Card>

            <Card>
              <CardBody>
                <Stat>
                  <StatLabel>Created</StatLabel>
                  <StatNumber fontSize="lg">{formatDate(urlStats.createdAt)}</StatNumber>
                  <StatHelpText>Date created</StatHelpText>
                </Stat>
              </CardBody>
            </Card>

            <Card>
              <CardBody>
                <Stat>
                  <StatLabel>Last Updated</StatLabel>
                  <StatNumber fontSize="lg">{formatDate(urlStats.updatedAt)}</StatNumber>
                  <StatHelpText>Last modified</StatHelpText>
                </Stat>
              </CardBody>
            </Card>
          </SimpleGrid>

          <Box>
            <Heading size="sm" mb={4}>URL Information</Heading>
            <SimpleGrid columns={{ base: 1, md: 2 }} spacing={4}>
              <Box>
                <Text fontWeight="bold">Short Code</Text>
                <Text>{urlStats.shortCode}</Text>
              </Box>
              <Box>
                <Text fontWeight="bold">ID</Text>
                <Text fontSize="sm" color="gray.600">{urlStats.id}</Text>
              </Box>
            </SimpleGrid>
          </Box>
        </CardBody>
      </Card>
    </Box>
  );
};

export default StatsPage; 