import logging

# Configure logging
logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(levelname)s - %(message)s')

class Logger:

    def __init__(self):
        self.logger = logging.getLogger(__name__)
    
    def debug(self, message):
        self.logger.debug(message)